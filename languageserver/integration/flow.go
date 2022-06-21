package integration

import (
	"github.com/onflow/cadence"
	"github.com/onflow/flow-cli/pkg/flowkit"
	"github.com/onflow/flow-cli/pkg/flowkit/services"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
)

type flowClient interface {
	ExecuteScript(
		code []byte,
		args []cadence.Value,
		scriptPath string,
		network string,
	) (cadence.Value, error)

	DeployContract(
		account *flowkit.Account,
		contractName string,
		contractSource []byte,
		update bool,
	) (*flow.Account, error)

	SendTransaction(
		signer *flowkit.Account,
		code []byte,
		codeFilename string,
		gasLimit uint64,
		args []cadence.Value,
		network string,
	) (*flow.Transaction, *flow.TransactionResult, error)

	GetAccount(address flow.Address) (*flow.Account, error)

	CreateAccount(
		signer *flowkit.Account,
		pubKeys []crypto.PublicKey,
		keyWeights []int,
		sigAlgo []crypto.SignatureAlgorithm,
		hashAlgo []crypto.HashAlgorithm,
		contractArgs []string,
	) (*flow.Account, error)
}

var _ flowClient = flowkitClient{}

type flowkitClient struct {
	services *services.Services
}

func (f flowkitClient) ExecuteScript(
	code []byte,
	args []cadence.Value,
	scriptPath string,
	network string,
) (cadence.Value, error) {
	return f.services.Scripts.Execute(code, args, scriptPath, network)
}

func (f flowkitClient) DeployContract(
	account *flowkit.Account,
	contractName string,
	contractSource []byte,
	update bool,
) (*flow.Account, error) {
	return f.services.Accounts.AddContract(account, contractName, contractSource, update)
}

func (f flowkitClient) SendTransaction(
	signer *flowkit.Account,
	code []byte,
	codeFilename string,
	gasLimit uint64,
	args []cadence.Value,
	network string,
) (*flow.Transaction, *flow.TransactionResult, error) {
	return f.services.Transactions.Send(signer, code, codeFilename, gasLimit, args, network)
}

func (f flowkitClient) GetAccount(address flow.Address) (*flow.Account, error) {
	return f.services.Accounts.Get(address)
}

func (f flowkitClient) CreateAccount(
	signer *flowkit.Account,
	pubKeys []crypto.PublicKey,
	keyWeights []int,
	sigAlgo []crypto.SignatureAlgorithm,
	hashAlgo []crypto.HashAlgorithm,
	contractArgs []string,
) (*flow.Account, error) {
	return f.services.Accounts.Create(signer, pubKeys, keyWeights, sigAlgo, hashAlgo, contractArgs)
}
