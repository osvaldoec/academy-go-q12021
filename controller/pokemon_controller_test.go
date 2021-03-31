package controller

import (
	"errors"
	"pokemon/controller/mocks"
	"pokemon/entity"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

var (
	validResponse = &entity.Pokemon{
		ID:             1,
		Name:           "getByIDTest",
		BaseExperience: 50,
	}
	validInstertResponse = &entity.Pokemon{
		ID:             4,
		Name:           "insertResponseTest",
		BaseExperience: 50,
	}
)

//mockgen -source=controller/pokemon_controller.go -destination=controller/mocks/pokemon_controller.go -package=mocks
func Test_PokemonController(t *testing.T) {
	tests := []struct {
		name                    string
		expectedParams          string
		expectedUsecaseResponse *entity.Pokemon
		expectUsecaseCall       bool
		expectedError           error
		wantError               bool
	}{
		{
			name:                    "OK, GetByID",
			expectedParams:          "1",
			expectedUsecaseResponse: validResponse,
			expectUsecaseCall:       true,
			wantError:               false,
			expectedError:           nil,
		},
		{
			name:                    "Not Found, GetByID",
			expectedParams:          "100000",
			expectedUsecaseResponse: nil,
			expectUsecaseCall:       true,
			wantError:               true,
			expectedError:           errors.New("Not found"),
		},
		{
			name:                    "OK, InsertByID",
			expectedParams:          "4",
			expectedUsecaseResponse: validInstertResponse,
			expectUsecaseCall:       true,
			wantError:               false,
			expectedError:           nil,
		},
		{
			name:                    "Not Found, InsertByID",
			expectedParams:          "10000000",
			expectedUsecaseResponse: nil,
			expectUsecaseCall:       true,
			wantError:               true,
			expectedError:           errors.New("Not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			u := mocks.NewMockUseCase(mockCtrl)

			if tt.expectUsecaseCall {
				u.EXPECT().GetByID(tt.expectedParams).Return(tt.expectedUsecaseResponse, tt.expectedError)
			}

			c := New(u)
			response, err := c.UseCase.GetByID(tt.expectedParams)
			assert.Equal(t, response, tt.expectedUsecaseResponse)

			if tt.wantError {
				assert.NotNil(t, err)
				assert.Equal(t, err, tt.expectedError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
