package utils

import (
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1Native"
)

//TableName return the full table name taken into account the parent
func TableName(t corev1.TableReader) (value string) {
	pid := corev1Native.NewIDReader([]byte(t.Meta().Application()+"/"+t.Meta().Parent()), false)
	if pid != nil {
		return pid.ID() + t.Meta().Name()

	}
	return t.Meta().Name()
}

func TableNameFromID(pid corev1.ResourceIDReader) (value string) {
	names := []string{}
	var obj corev1.ResourceIDReader
	if pid == nil {
		return
	}

	obj = pid
	for obj != nil {
		if obj.Kind() != "Table" && obj.Kind() != "Resource" {
			return
		}
		names = append(names, obj.ID())
		obj = obj.Parent()

	}

	for i := 0; i < len(names); i++ {
		value = value + names[len(names)-1-i]
	}

	return
}