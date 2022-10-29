package generated

import (
	"github.com/gowcar/gboot-sandbox/internal/user"
	"github.com/gowcar/gboot/pkg/annotation"
)

type __USER__PKG__ANNOTATION__ struct{}

func (__USER__PKG__ANNOTATION__) PackageName() string {
	return "user"
}
func (__USER__PKG__ANNOTATION__) AllAnnotations() []annotation.Annotation {
	annotations := make([]annotation.Annotation, 0)
	var anno annotation.Annotation

	anno = annotation.Annotation{
		PackageName:      "user",
		AnnotationName:   "@Config",
		AnnotationTarget: annotation.Var,
		TargetTypeName:   "",
		TargetName:       "user.DefaultUserName",
		TargetObject:     &user.DefaultUserName,
		RawData:          "\"user.default_username\"",
		Params: map[string]any{
			"default": "user.default_username",
		},
	}
	annotations = append(annotations, anno)

	anno = annotation.Annotation{
		PackageName:      "user",
		AnnotationName:   "@Config",
		AnnotationTarget: annotation.Var,
		TargetTypeName:   "",
		TargetName:       "user.Timeout",
		TargetObject:     &user.Timeout,
		RawData:          "\"application.timeout\"",
		Params: map[string]any{
			"default": "application.timeout",
		},
	}
	annotations = append(annotations, anno)

	anno = annotation.Annotation{
		PackageName:      "user",
		AnnotationName:   "@GetMapping",
		AnnotationTarget: annotation.Func,
		TargetTypeName:   "",
		TargetName:       "user.GlobalController",
		TargetObject:     user.GlobalController,
		RawData:          "\"/api/static/config\"",
		Params: map[string]any{
			"default": "/api/static/config",
		},
	}
	annotations = append(annotations, anno)

	anno = annotation.Annotation{
		PackageName:      "user",
		AnnotationName:   "@RequestMapping",
		AnnotationTarget: annotation.Struct,
		TargetTypeName:   "",
		TargetName:       "user.UserController",
		TargetObject:     &user.UserController{},
		RawData:          "\"/api/users\"",
		Params: map[string]any{
			"default": "/api/users",
		},
	}
	annotations = append(annotations, anno)

	anno = annotation.Annotation{
		PackageName:      "user",
		AnnotationName:   "@RestController",
		AnnotationTarget: annotation.Struct,
		TargetTypeName:   "",
		TargetName:       "user.UserController",
		TargetObject:     &user.UserController{},
		RawData:          "",
		Params: map[string]any{
			"default": "",
		},
	}
	annotations = append(annotations, anno)

	anno = annotation.Annotation{
		PackageName:      "user",
		AnnotationName:   "@Config",
		AnnotationTarget: annotation.StructField,
		TargetTypeName:   "user.UserController",
		TargetName:       "FetchSize",
		TargetObject:     nil,
		RawData:          "\"user.fetch_size\"",
		Params: map[string]any{
			"default": "user.fetch_size",
		},
	}
	annotations = append(annotations, anno)

	anno = annotation.Annotation{
		PackageName:      "user",
		AnnotationName:   "@GetMapping",
		AnnotationTarget: annotation.StructMethod,
		TargetTypeName:   "user.UserController",
		TargetName:       "UserHandler",
		TargetObject:     nil,
		RawData:          "\"/\"",
		Params: map[string]any{
			"default": "/",
		},
	}
	annotations = append(annotations, anno)

	return annotations

}

