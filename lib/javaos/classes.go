package javaos

type JavaClassReader interface {
	Name() string
	Read(s *Stream, cd *ClassDesc) interface{}
}

var javaClasses = []JavaClassReader{
	&java_util_Date{},
	&java_util_HashMap{},
}

func javaClassReaderFor(desc *ClassDesc) JavaClassReader {
	for _, jc := range javaClasses {
		if jc.Name() == desc.Name {
			return jc
		}
	}
	return nil
}
