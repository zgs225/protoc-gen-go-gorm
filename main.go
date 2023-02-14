package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/zgs225/protoc-gen-go-gorm/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	modelSuffix string

	timePackage = protogen.GoImportPath("time")
	sqlPackage  = protogen.GoImportPath("database/sql")
)

func main() {
	flag := flag.FlagSet{}

	flag.StringVar(&modelSuffix, "suffix", "Model", "模型名称后缀")

	opts := protogen.Options{
		ParamFunc: flag.Set,
	}

	opts.Run(func(p *protogen.Plugin) error {
		for _, f := range p.Files {
			if !f.Generate {
				continue
			}
			generateFile(p, f)
		}
		return nil
	})
}

func generateFile(p *protogen.Plugin, f *protogen.File) {
	g := p.NewGeneratedFile(f.GeneratedFilenamePrefix+".gorm.go", f.GoImportPath)
	g.P("// Code generated by protoc-gen-go-gorm. DO NOT EDIT.")
	g.P()
	g.P("package ", f.GoPackageName)
	g.P()

	for _, msg := range f.Messages {
		opts := msg.Desc.Options().(*descriptorpb.MessageOptions)
		has := opts.ProtoReflect().Has(proto.E_Model.TypeDescriptor())
		if !has {
			continue
		}
		ext := proto.E_Model.InterfaceOf(opts.ProtoReflect().Get(proto.E_Model.TypeDescriptor())).(*proto.Model)
		generateMessage(p, g, f, msg, ext)
	}
}

func generateMessage(p *protogen.Plugin, g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message, ext *proto.Model) {
	g.P(m.Comments.Leading)
	// 生成结构体
	g.P("type ", modelName(m), " struct {")
	for _, field := range m.Fields {
		fieldName := field.GoName
		fieldType, _ := fieldGoType(g, field)

		g.P(fieldName, " ", fieldType, " ", gormFieldTag(field))
	}
	g.P("}")

	if len(ext.GetTableName()) > 0 {
		g.P("func(m ", modelName(m), ") TableName() string {")
		g.P("return \"", ext.GetTableName(), "\"")
		g.P("}")
	}
}

func modelName(m *protogen.Message) string {
	return m.GoIdent.GoName + modelSuffix
}

func gormFieldTag(f *protogen.Field) string {
	has := f.Desc.Options().ProtoReflect().Has(proto.E_Field.TypeDescriptor())
	if !has {
		return ""
	}
	ext, ok := proto.E_Field.InterfaceOf(f.Desc.Options().ProtoReflect().Get(proto.E_Field.TypeDescriptor())).(*proto.Field)
	if !ok {
		return ""
	}

	parts := make([]string, 0)

	if len(ext.GetColumn()) > 0 {
		parts = append(parts, "column:"+ext.GetColumn())
	}

	if len(ext.GetType()) > 0 {
		parts = append(parts, "type:"+ext.GetType())
	}

	if len(ext.GetSerializer()) > 0 {
		parts = append(parts, "serializer:"+ext.GetSerializer())
	}

	if ext.GetSize() > 0 {
		parts = append(parts, "size:", strconv.FormatInt(ext.GetSize(), 10))
	}

	if ext.GetPk() {
		parts = append(parts, "primaryKey")
	}

	if ext.GetUnique() {
		parts = append(parts, "unique")
	}

	if len(ext.GetDefaultValue()) > 0 {
		parts = append(parts, "default:"+ext.GetDefaultValue())
	}

	if ext.GetPrecision() > 0 {
		parts = append(parts, "precision:"+strconv.FormatInt(int64(ext.GetPrecision()), 10))
	}

	if ext.GetScale() > 0 {
		parts = append(parts, "scale:"+strconv.FormatInt(int64(ext.GetScale()), 10))
	}

	if ext.GetNotNull() {
		parts = append(parts, "not null")
	}

	if ext.GetAutoIncrement() {
		parts = append(parts, "autoIncrement")
	}

	if ext.GetAutoIncrementStep() > 0 {
		parts = append(parts, "autoIncrementIncrement:"+strconv.FormatInt(int64(ext.GetAutoIncrementStep()), 10))
	}

	if ext.GetEmbedded() {
		parts = append(parts, "embedded")
	}

	if len(ext.GetEmbeddedPrefix()) > 0 {
		parts = append(parts, "embeddedPrefix:"+ext.GetEmbeddedPrefix())
	}

	switch ext.GetAutoCreateTime() {
	case proto.AutoTimeUnit_AutoTimeUnitMilli:
		parts = append(parts, "autoCreateTime:milli")
	case proto.AutoTimeUnit_AutoTimeUnitNano:
		parts = append(parts, "autoCreateTime:nano")
	case proto.AutoTimeUnit_AutoTimeUnitSecond:
		parts = append(parts, "autoCreateTime")
	}

	switch ext.GetAutoUpdateTime() {
	case proto.AutoTimeUnit_AutoTimeUnitMilli:
		parts = append(parts, "autoUpdateTime:milli")
	case proto.AutoTimeUnit_AutoTimeUnitNano:
		parts = append(parts, "autoUpdateTime:nano")
	case proto.AutoTimeUnit_AutoTimeUnitSecond:
		parts = append(parts, "autoUpdateTime")
	}

	if len(ext.GetIndex()) > 0 {
		parts = append(parts, "index:"+ext.GetIndex())
	}

	if len(ext.GetUniqIndex()) > 0 {
		parts = append(parts, "uniqueIndex:"+ext.GetUniqIndex())
	}

	if len(ext.GetCheck()) > 0 {
		parts = append(parts, "check:", ext.GetCheck())
	}

	switch ext.GetWritePerm() {
	case proto.WritePermission_WritePermissionNone:
		parts = append(parts, "<-:false")
	case proto.WritePermission_WritePermissionCreate:
		parts = append(parts, "<-:create")
	case proto.WritePermission_WritePermissionUpdate:
		parts = append(parts, "<-:update")
	}

	if ext.GetNoReadPerm() {
		parts = append(parts, "->:false")
	}

	if len(ext.GetIgnore()) > 0 {
		parts = append(parts, ext.GetIgnore())
	}

	if len(parts) == 0 {
		return ""
	}

	return fmt.Sprintf("`gorm:\"%s\"`", strings.Join(parts, ";"))
}

func fieldGoType(g *protogen.GeneratedFile, f *protogen.Field) (string, bool) {
	if f.Desc.IsWeak() {
		return "struct{}", false
	}

	pointer := f.Desc.HasPresence()
	goType := ""

	opt, ok := proto.E_Field.InterfaceOf(f.Desc.Options().ProtoReflect().Get(proto.E_Field.TypeDescriptor())).(*proto.Field)

	// 使用自定义类型
	if ok && opt.GetExtType() != nil {
		pkg := protogen.GoImportPath(opt.GetExtType().GetImportPath())
		goType = g.QualifiedGoIdent(pkg.Ident(opt.GetExtType().GetType()))
		return goType, false
	}

	switch f.Desc.Kind() {
	case protoreflect.BoolKind:
		if ok && opt.GetNullableType() {
			goType = g.QualifiedGoIdent(sqlPackage.Ident("NullBool"))
		} else {
			goType = "bool"
		}
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(f.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if ok && opt.GetNullableType() {
			goType = g.QualifiedGoIdent(sqlPackage.Ident("NullInt64"))
		} else {
			goType = "int32"
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if ok && opt.GetNullableType() {
			goType = g.QualifiedGoIdent(sqlPackage.Ident("NullInt64"))
		} else {
			goType = "uint32"
		}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if ok && opt.GetNullableType() {
			goType = g.QualifiedGoIdent(sqlPackage.Ident("NullInt64"))
		} else {
			goType = "int64"
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if ok && opt.GetNullableType() {
			goType = g.QualifiedGoIdent(sqlPackage.Ident("NullInt64"))
		} else {
			goType = "uint64"
		}
	case protoreflect.FloatKind:
		if ok && opt.GetNullableType() {
			goType = g.QualifiedGoIdent(sqlPackage.Ident("NullFloat64"))
		} else {
			goType = "float32"
		}
	case protoreflect.DoubleKind:
		if ok && opt.GetNullableType() {
			goType = g.QualifiedGoIdent(sqlPackage.Ident("NullFloat64"))
		} else {
			goType = "float64"
		}
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false
	case protoreflect.StringKind:
		if ok && opt.GetNullableType() {
			goType = g.QualifiedGoIdent(sqlPackage.Ident("NullString"))
		} else {
			goType = "string"
		}
	case protoreflect.MessageKind, protoreflect.GroupKind:
		if f.Message.GoIdent.GoImportPath == "github.com/golang/protobuf/ptypes/timestamp" {
			if ok && opt.GetNullableType() {
				goType = g.QualifiedGoIdent(sqlPackage.Ident("NullTime"))
			} else {
				goType = g.QualifiedGoIdent(timePackage.Ident("Time"))
			}
		} else {
			goType = "*" + g.QualifiedGoIdent(f.Message.GoIdent)
		}

		pointer = false
	}

	switch {
	case f.Desc.IsList():
		goType = "[]" + goType
		pointer = false
	case f.Desc.IsMap():
		keyType, _ := fieldGoType(g, f.Message.Fields[0])
		valType, _ := fieldGoType(g, f.Message.Fields[1])
		goType = fmt.Sprintf("map[%s]%s", keyType, valType)
		pointer = false
	}

	return goType, pointer
}
