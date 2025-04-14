package stores

import (
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type FeatureFlagEnvironment struct {
	Name    string `hcl:",label"`
	Enabled bool   `hcl:"enabled"`
}

type FeatureFlag struct {
	Name string `hcl:",label"`

	TeamResponsible string `hcl:"team_responsible"`
	ContactPerson   string `hcl:"contact_person"`

	Environments []FeatureFlagEnvironment `hcl:"environment,block"`
}

type FeatureFlags struct {
	FeatureFlags []FeatureFlag `hcl:"feature_flag,block"`
}

type FeatureFlagsStore struct {
	FeatureFlags FeatureFlags
	Lookup       map[string]map[string]bool
}

func ReadFeatureFlags(path string) FeatureFlagsStore {
	// Decode basic feature flags
	var featureFlags FeatureFlags
	err := hclsimple.DecodeFile(path, nil, &featureFlags)
	if err != nil {
		log.Fatalf("Failed to load feature flags: %s", err)
	}

	// Create fast lookup
	lookup := make(map[string]map[string]bool)
	for _, featureFlag := range featureFlags.FeatureFlags {
		lookupHere, ok := lookup[featureFlag.Name]
		if !ok {
			lookupHere = make(map[string]bool)
			lookup[featureFlag.Name] = lookupHere
		}

		for _, featureFlagEnvironment := range featureFlag.Environments {
			lookupHere[featureFlagEnvironment.Name] = featureFlagEnvironment.Enabled
		}
	}

	return FeatureFlagsStore{FeatureFlags: featureFlags, Lookup: lookup}
}

func (featureFlags FeatureFlags) Find(featureFlagName string) (FeatureFlag, bool) {
	var featureFlag FeatureFlag
	found := false

	for _, f := range featureFlags.FeatureFlags {
		if f.Name == featureFlagName {
			featureFlag = f
			found = true
			break
		}
	}

	if !found {
		return FeatureFlag{}, false
	}

	return featureFlag, true
}

func (featureFlag FeatureFlag) Find(environmentName string) (FeatureFlagEnvironment, bool) {
	var featureFlagEnvironment FeatureFlagEnvironment
	found := false

	for _, e := range featureFlag.Environments {
		if e.Name == environmentName {
			featureFlagEnvironment = e
			found = true
			break
		}
	}

	if !found {
		return FeatureFlagEnvironment{}, false
	}

	return featureFlagEnvironment, true
}
