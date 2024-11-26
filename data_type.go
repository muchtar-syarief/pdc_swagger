package pdc_swagger

import "reflect"

type DataType string

const (
	DataTypeNumber  DataType = "number"
	DataTypeInteger DataType = "integer"
	DataTypeString  DataType = "string"
	DataTypeBoolean DataType = "boolean"
	DataTypeArray   DataType = "array"
	DataTypeObject  DataType = "object"
	DataTypeUnknown DataType = "unknown"
)

func GetDataTypeMapper(dataType reflect.Kind) DataType {
	switch dataType {
	case reflect.String:
		return DataTypeString
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return DataTypeInteger
	case reflect.Float32, reflect.Float64:
		return DataTypeNumber
	case reflect.Array, reflect.Slice:
		return DataTypeArray
	case reflect.Map, reflect.Pointer, reflect.Struct, reflect.Interface:
		return DataTypeObject
	}

	return DataTypeUnknown
}

type DataFormat string

const (
	DataFormatInt32    DataFormat = "int32"
	DataFormatInt64    DataFormat = "int64"
	DataFormatFloat    DataFormat = "float"
	DataFormatDouble   DataFormat = "double"
	DataFormatPassword DataFormat = "password"
	DataFormatEmail    DataFormat = "email"
	DataFormatUUID     DataFormat = "uuid"
	DataFormatZipCode  DataFormat = "zip-code"
)

var DataTypeFormatMap = map[DataType][]DataFormat{
	DataTypeString:  {DataFormatPassword, DataFormatEmail, DataFormatUUID, DataFormatZipCode},
	DataTypeInteger: {DataFormatInt32, DataFormatInt64},
	DataTypeNumber:  {DataFormatFloat, DataFormatDouble},
}
