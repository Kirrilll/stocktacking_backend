package fixtures

import (
	"encoding/json"
	"io"
	"os"
	"stocktacking_backend/internal/domain/entity/branch_entity"
	"stocktacking_backend/internal/domain/entity/category_entity"
	"stocktacking_backend/internal/domain/entity/item_entity"
	"stocktacking_backend/internal/domain/entity/organization_entity"
	"stocktacking_backend/internal/domain/entity/report"
	"stocktacking_backend/internal/domain/entity/storage_entity"
	"stocktacking_backend/internal/domain/entity/user_entity"
	"stocktacking_backend/pkg/configurations/plugin/entity"
)

type BranchFixture struct{}

func (BranchFixture) GetEntity() any {
	return &branch_entity.Branch{}
}

func (f BranchFixture) FixturesData() (any, error) {
	jsonFile, err := os.Open("fixtures_data/entities/branches.json ")
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data []branch_entity.Branch
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

type CategoryFixture struct{}

func (CategoryFixture) GetEntity() any {
	return &category_entity.Category{}
}

func (f CategoryFixture) FixturesData() (any, error) {
	jsonFile, err := os.Open("fixtures_data/entities/categoties.json")
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data []category_entity.Category
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

type ItemFixture struct{}

func (ItemFixture) GetEntity() any {
	return &item_entity.Item{}
}

func (f ItemFixture) FixturesData() (any, error) {
	jsonFile, err := os.Open("fixtures_data/entities/items.json")
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data []item_entity.Item
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

type OrganizationFixture struct{}

func (OrganizationFixture) GetEntity() any {
	return &organization_entity.Organization{}
}

func (f OrganizationFixture) FixturesData() (any, error) {
	jsonFile, err := os.Open("fixtures_data/entities/organization.json")
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data []organization_entity.Organization
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

type ReportFixture struct{}

func (ReportFixture) GetEntity() any {
	return &report.Report{}
}

func (f ReportFixture) FixturesData() (any, error) {
	jsonFile, err := os.Open("fixtures_data/entities/reports.json")
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data []report.Report
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

type StorageFixture struct{}

func (StorageFixture) GetEntity() any {
	return &storage_entity.Storage{}
}

func (f StorageFixture) FixturesData() (any, error) {
	jsonFile, err := os.Open("fixtures_data/product/carts.json")
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data []storage_entity.Storage
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

type UserFixture struct{}

func (UserFixture) GetEntity() any {
	return &user_entity.User{}
}

func (f UserFixture) FixturesData() (any, error) {
	jsonFile, err := os.Open("fixtures_data/entities/users.json")
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data []user_entity.User
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

type ConfigurationFixture struct{}

func (ConfigurationFixture) GetEntity() any {
	return &entity.Configuration{}
}

func (f ConfigurationFixture) FixturesData() (any, error) {
	jsonFile, err := os.Open("fixtures_data/entities/conf.json")
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data []entity.Configuration
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

type OptionFixture struct{}

func (OptionFixture) GetEntity() any {
	return &entity.Option{}
}

func (f OptionFixture) FixturesData() (any, error) {
	jsonFile, err := os.Open("fixtures_data/entities/conf.json")
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data []entity.Option
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}
