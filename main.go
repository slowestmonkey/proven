package main

import (
	profileRepository "proven/adapters/profile"
	profile "proven/core/profile"
)

func main() {
	var connection string

	profileRepo := profileRepository.NewProfileRepository(connection)
	profileService := profile.NewProfileService(profileRepo)

}
