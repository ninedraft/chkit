package image

import (
	"fmt"

	"github.com/containerum/chkit/pkg/chkitErrors"
	"github.com/containerum/chkit/pkg/util/validation"
	"github.com/containerum/kube-client/pkg/model"
)

const (
	ErrInvalidUpdateImage chkitErrors.Err = "invalid update image"
)

func ValidateImage(img model.UpdateImage) error {
	var errs []error
	if err := validation.ValidateImageName(img.Image); err != nil {
		errs = append(errs, fmt.Errorf("\n + invalid image %q", img.Image))
	}
	if err := validation.ValidateLabel(img.Container); err != nil {
		errs = append(errs, fmt.Errorf("\n + invalid container label %q", img.Container))
	}
	if len(errs) == 0 {
		return nil
	}
	return ErrInvalidUpdateImage.Wrap(errs...)
}
