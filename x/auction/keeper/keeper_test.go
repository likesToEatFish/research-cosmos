package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/DongCoNY/research-cosmos/app/apptesting"
	"github.com/DongCoNY/research-cosmos/x/auction/types"
)

type KeeperTestSuite struct {
	apptesting.AppTestHelper
	suite.Suite
	queryClient types.QueryClient
}

// Test helpers
func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.queryClient = types.NewQueryClient(s.QueryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
