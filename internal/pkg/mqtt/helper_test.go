package mqtt

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type MqttTestSuite struct {
	suite.Suite
}

func (suite *MqttTestSuite) SetupTest() {
}

func (suite *MqttTestSuite) TestGetIdsFromTopic() {
	expectedIds := []string{"examplePlugin"}
	ids, err := GetIdsFromTopic("cmd-manager/examplePlugin/execute", "cmd-manager/+/execute")
	suite.Require().NoError(err)
	suite.Require().Equal(expectedIds, ids)

	ids, err = GetIdsFromTopic("cmd-manager/execute", "cmd-manager/+/execute")
	suite.Require().Error(err)

	ids, err = GetIdsFromTopic("ploogin/examplePlugin/execute", "cmd-manager/+/execute")
	suite.Require().Error(err)

	ids, err = GetIdsFromTopic("ploogin/examplePlugin/execute", "cmd-manager/execute")
	suite.Require().Error(err)

	ids, err = GetIdsFromTopic("cmd-manager/examplePlugin/execute", "cmd-manager/examplePlugin/execute")
	suite.Require().Error(err)

	ids, err = GetIdsFromTopic("cmd-manager/examplePlugin/execute/example2/abc", "cmd-manager/+/execute/+/abc")
	suite.Require().NoError(err)
	suite.Require().Equal([]string{"examplePlugin", "example2"}, ids)
}

func (suite *MqttTestSuite) TestCreateTopicWithIds() {
	ids, err := CreateTopicWithIds("cmd-manager/+/execute", "exampleId")
	suite.Require().NoError(err)
	suite.Require().Equal("cmd-manager/exampleId/execute", ids)

	ids, err = CreateTopicWithIds("cmd-manager/+/execute/+/", "exampleId1", "exampleId2")
	suite.Require().NoError(err)
	suite.Require().Equal("cmd-manager/exampleId1/execute/exampleId1/", ids)

	ids, err = CreateTopicWithIds("cmd-manager/+/execute/+/", "exampleId")
	suite.Require().Error(err)

	ids, err = CreateTopicWithIds("cmd-manager/+/execute/+/", "exampleId", "")
	suite.Require().Error(err)
}

func TestGetIdsFromTopic(t *testing.T) {
	suite.Run(t, new(MqttTestSuite))
}
