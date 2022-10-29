package generated

import (
	"github.com/gowcar/gboot-sandbox/internal/user"
	"github.com/gowcar/gboot/pkg/annotation"
)

type __USER__PKG__ANNOTATION__ struct{}

func (__USER__PKG__ANNOTATION__) PackageName() string {
	return "user"
}

func (__USER__PKG__ANNOTATION__) PackageVariableAnnotations() []annotation.Annotation {
	annotations := make([]annotation.Annotation, 0)
	var anno annotation.Annotation

	anno = annotation.Annotation{
		PackageName:    "user",
		AnnotationName: "@Config",
		TargetName:     "user.DefaultUserName",
		TargetObject:   &user.DefaultUserName,
		Params: map[string]any{
			"default": "user.default_username",
		},
		RawData: "\"user.default_username\"",
	}

	annotations = append(annotations, anno)

	anno = annotation.Annotation{
		PackageName:    "user",
		AnnotationName: "@Config",
		TargetName:     "user.Timeout",
		TargetObject:   &user.Timeout,
		Params: map[string]any{
			"default": "application.timeout",
		},
		RawData: "\"application.timeout\"",
	}
	annotations = append(annotations, anno)
	return annotations
}

func (__USER__PKG__ANNOTATION__) PackageFunctionAnnotations() []annotation.Annotation {
	return nil
}

func (__USER__PKG__ANNOTATION__) StructAnnotations() []annotation.Annotation {
	return nil
}

func (__USER__PKG__ANNOTATION__) StructFieldAnnotations() []annotation.Annotation {
	return nil
}

func (__USER__PKG__ANNOTATION__) StructMethodAnnotations() []annotation.Annotation {
	return nil
}
