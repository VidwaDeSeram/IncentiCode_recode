package actions

import (
	"github.com/VidwaDeSeram/IncentiCode_recode/entities"
	"github.com/VidwaDeSeram/IncentiCode_recode/stepper"
)

func InstallRecode(
	stepper stepper.Stepper,
	cloudService entities.CloudService,
	recodeConfig *entities.Config,
) error {

	err := cloudService.CreateRecodeConfigStorage(stepper)

	if err != nil {
		return err
	}

	return cloudService.SaveRecodeConfig(
		stepper,
		recodeConfig,
	)
}
