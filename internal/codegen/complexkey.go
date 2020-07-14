package codegen

import (
	. "github.com/dave/jennifer/jen"
)

type ComplexKey struct {
	NamedType
	Key    Identifier
	Params Identifier
}

func (ck *ComplexKey) InnerTypes() IdentifierSet {
	return IdentifierSet{
		ck.Key:    true,
		ck.Params: true,
	}
}

func (ck *ComplexKey) GenerateCode() *Statement {
	def := AddWordWrappedComment(Empty(), ck.Doc).Line().
		Type().Id(ck.Name).
		StructFunc(func(def *Group) {
			def.Add(ck.Key.Qual())
			def.Id("Params").Op("*").Add(ck.Params.Qual()).Tag(JsonFieldTag("$params", false))
		}).Line().Line()

	record := &Record{
		NamedType: ck.NamedType,
		Fields:    TypeRegistry.Resolve(ck.Key).(*Record).Fields,
	}

	return AddRestLiEncode(def, record.Receiver(), ck.Name, func(def *Group) {
		record.unionFieldValidator(def)
		def.Line()
		record.generateEncoder(def, nil, &ck.Params)
		def.Return(Nil())
	})
}
