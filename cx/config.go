package base

import (
	"os"
)

//DbgGolangStackTrace ...
const DbgGolangStackTrace = true

//PROGRAM global reference to our program
var PROGRAM *CXProgram

//CXPATH ...
var CXPATH = os.Getenv("CXPATH") + "/"

//BINPATH ...
var BINPATH = CXPATH + "bin/"

//PKGPATH ...
var PKGPATH = CXPATH + "pkg/"

//SRCPATH ...
var SRCPATH = CXPATH + "src/"

//COREPATH ...
var COREPATH string

// nolint golint
const (
	StackOverflowError = "stack overflow"
	HeapExhaustedError = "heap exhausted"
	MAIN_FUNC          = "main"
	SYS_INIT_FUNC      = "*init"
	MAIN_PKG           = "main"
	OS_PKG             = "os"
	OS_ARGS            = "Args"

	NON_ASSIGN_PREFIX = "nonAssign"
	LOCAL_PREFIX      = "*lcl"
	LABEL_PREFIX      = "*lbl"

	// const CORE_MODULE = "core"
	ID_FN   = "identity"
	INIT_FN = "initDef"

	I32_SIZE = 4
	I64_SIZE = 8
	STR_SIZE = 4

	MARK_SIZE                = 1
	OBJECT_HEADER_SIZE       = 9
	OBJECT_GC_HEADER_SIZE    = 5
	FORWARDING_ADDRESS_SIZE  = 4
	OBJECT_SIZE              = 4
	CALLSTACK_SIZE           = 1000
	STACK_SIZE               = 1048576  // 1 Mb
	INIT_HEAP_SIZE           = 2097152  // 2 Mb
	MAX_HEAP_SIZE            = 67108864 // 64 Mb
	MIN_HEAP_FREE_RATIO      = 40
	MAX_HEAP_FREE_RATIO      = 70
	NULL_ADDRESS             = STACK_SIZE
	NULL_HEAP_ADDRESS_OFFSET = 4
	NULL_HEAP_ADDRESS        = 0
	STR_HEADER_SIZE          = 4
	TYPE_POINTER_SIZE        = 4
	SLICE_HEADER_SIZE        = 8
	MAX_UINT32               = ^uint32(0)
	MIN_UINT32               = 0
	MAX_INT32                = int(MAX_UINT32 >> 1)
	MIN_INT32                = -MAX_INT32 - 1
)

// nolint golint
var (
	MEMORY_SIZE = STACK_SIZE + INIT_HEAP_SIZE + TYPE_POINTER_SIZE
	BASIC_TYPES = []string{
		"bool", "str", "byte", "i32", "i64", "f32", "f64",
		"[]bool", "[]str", "[]byte", "[]i32", "[]i64", "[]f32", "[]f64",
	}
)

// nolint golint
const (
	CX_SUCCESS = iota
	CX_RUNTIME_ERROR
	CX_PANIC // 2
	CX_COMPILATION_ERROR
	CX_INTERNAL_ERROR
	CX_ASSERT
)

// nolint golint
const (
	DECL_POINTER = iota // 0
	DECL_ARRAY          // 1
	DECL_SLICE          // 2
	DECL_STRUCT         // 3
	DECL_BASIC          // 4
)

// what to write
// nolint golint
const (
	PASSBY_VALUE = iota
	PASSBY_REFERENCE
)

// nolint golint
const (
	DEREF_ARRAY = iota
	DEREF_FIELD
	DEREF_POINTER
	DEREF_DEREF
)

// nolint golint
const (
	TYPE_UNDEFINED = iota
	TYPE_AFF
	TYPE_BOOL
	TYPE_BYTE
	TYPE_STR
	TYPE_F32
	TYPE_F64
	TYPE_I8
	TYPE_I16
	TYPE_I32
	TYPE_I64
	TYPE_UI8
	TYPE_UI16
	TYPE_UI32
	TYPE_UI64

	TYPE_THRESHOLD

	TYPE_CUSTOM
	TYPE_POINTER
	TYPE_IDENTIFIER
)

// TypeCounter ...
var TypeCounter int

// TypeCodes ..
var TypeCodes = map[string]int{
	"ident": TYPE_IDENTIFIER,
	"aff":   TYPE_AFF,
	"bool":  TYPE_BOOL,
	"byte":  TYPE_BYTE,
	"str":   TYPE_STR,
	"f32":   TYPE_F32,
	"f64":   TYPE_F64,
	"i8":    TYPE_I8,
	"i16":   TYPE_I16,
	"i32":   TYPE_I32,
	"i64":   TYPE_I64,
	"ui8":   TYPE_UI8,
	"ui16":  TYPE_UI16,
	"ui32":  TYPE_UI32,
	"ui64":  TYPE_UI64,
	"und":   TYPE_UNDEFINED,
}

//TypeNames ...
var TypeNames = map[int]string{
	TYPE_IDENTIFIER: "ident",
	TYPE_AFF:        "aff",
	TYPE_BOOL:       "bool",
	TYPE_BYTE:       "byte",
	TYPE_STR:        "str",
	TYPE_F32:        "f32",
	TYPE_F64:        "f64",
	TYPE_I8:         "i8",
	TYPE_I16:        "i16",
	TYPE_I32:        "i32",
	TYPE_I64:        "i64",
	TYPE_UI8:        "ui8",
	TYPE_UI16:       "ui16",
	TYPE_UI32:       "ui32",
	TYPE_UI64:       "ui64",
	TYPE_UNDEFINED:  "und",
}

// memory locations
// nolint golint
const (
	MEM_STACK = iota
	MEM_HEAP
	MEM_DATA
)
