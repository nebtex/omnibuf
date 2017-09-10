// Code generated by mockery v1.0.0
package mocks

import hybrids "github.com/nebtex/hybrids/golang/hybrids"
import mock "github.com/stretchr/testify/mock"
import oreflection "github.com/nebtex/omniql/commons/golang/oreflection"

// OType is an autogenerated mock type for the OType type
type OType struct {
	mock.Mock
}

// Application provides a mock function with given fields:
func (_m *OType) Application() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Enumeration provides a mock function with given fields:
func (_m *OType) Enumeration() hybrids.TableReader {
	ret := _m.Called()

	var r0 hybrids.TableReader
	if rf, ok := ret.Get(0).(func() hybrids.TableReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hybrids.TableReader)
		}
	}

	return r0
}

// FieldCount provides a mock function with given fields:
func (_m *OType) FieldCount() hybrids.FieldNumber {
	ret := _m.Called()

	var r0 hybrids.FieldNumber
	if rf, ok := ret.Get(0).(func() hybrids.FieldNumber); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(hybrids.FieldNumber)
	}

	return r0
}

// HybridType provides a mock function with given fields:
func (_m *OType) HybridType() hybrids.Types {
	ret := _m.Called()

	var r0 hybrids.Types
	if rf, ok := ret.Get(0).(func() hybrids.Types); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(hybrids.Types)
	}

	return r0
}

// Id provides a mock function with given fields:
func (_m *OType) Id() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Items provides a mock function with given fields:
func (_m *OType) Items() oreflection.OType {
	ret := _m.Called()

	var r0 oreflection.OType
	if rf, ok := ret.Get(0).(func() oreflection.OType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oreflection.OType)
		}
	}

	return r0
}

// Kind provides a mock function with given fields:
func (_m *OType) Kind() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// LookupEnumeration provides a mock function with given fields:
func (_m *OType) LookupEnumeration() oreflection.LookupEnumeration {
	ret := _m.Called()

	var r0 oreflection.LookupEnumeration
	if rf, ok := ret.Get(0).(func() oreflection.LookupEnumeration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oreflection.LookupEnumeration)
		}
	}

	return r0
}

// LookupFields provides a mock function with given fields:
func (_m *OType) LookupFields() oreflection.LookupFields {
	ret := _m.Called()

	var r0 oreflection.LookupFields
	if rf, ok := ret.Get(0).(func() oreflection.LookupFields); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oreflection.LookupFields)
		}
	}

	return r0
}

// LookupTableOnUnion provides a mock function with given fields:
func (_m *OType) LookupTableOnUnion() oreflection.LookupTableOnUnion {
	ret := _m.Called()

	var r0 oreflection.LookupTableOnUnion
	if rf, ok := ret.Get(0).(func() oreflection.LookupTableOnUnion); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oreflection.LookupTableOnUnion)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *OType) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Resource provides a mock function with given fields:
func (_m *OType) Resource() hybrids.TableReader {
	ret := _m.Called()

	var r0 hybrids.TableReader
	if rf, ok := ret.Get(0).(func() hybrids.TableReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hybrids.TableReader)
		}
	}

	return r0
}

// Table provides a mock function with given fields:
func (_m *OType) Table() hybrids.TableReader {
	ret := _m.Called()

	var r0 hybrids.TableReader
	if rf, ok := ret.Get(0).(func() hybrids.TableReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hybrids.TableReader)
		}
	}

	return r0
}

// Union provides a mock function with given fields:
func (_m *OType) Union() hybrids.TableReader {
	ret := _m.Called()

	var r0 hybrids.TableReader
	if rf, ok := ret.Get(0).(func() hybrids.TableReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hybrids.TableReader)
		}
	}

	return r0
}