package gateway

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/external"
)

// MockPaymentRepository implements repository.IPaymentRepository for testing
type MockPaymentRepository struct {
	CreateCalled        bool
	CreateInput         *model.Payment
	FindOneCalled       bool
	FindOneInput        string
	FindOneOutput       *model.Payment
	FindOneErr          error
	UpdateCalled        bool
	UpdateInput         *model.Payment
	FindByOrderIdCalled bool
	FindByOrderIdInput  string
	FindByOrderIdOutput *model.Payment
	FindByOrderIdErr    error
}

func (m *MockPaymentRepository) Create(ctx context.Context, payment *model.Payment) error {
	m.CreateCalled = true
	m.CreateInput = payment
	return nil
}
func (m *MockPaymentRepository) FindOne(ctx context.Context, id string) (*model.Payment, error) {
	m.FindOneCalled = true
	m.FindOneInput = id
	return m.FindOneOutput, m.FindOneErr
}
func (m *MockPaymentRepository) Update(ctx context.Context, payment *model.Payment) error {
	m.UpdateCalled = true
	m.UpdateInput = payment
	return nil
}
func (m *MockPaymentRepository) FindByOrderId(ctx context.Context, id string) (*model.Payment, error) {
	m.FindByOrderIdCalled = true
	m.FindByOrderIdInput = id
	return m.FindByOrderIdOutput, m.FindByOrderIdErr
}

// MockPaymentService implements external.IPaymentService for testing
type MockPaymentService struct{}

// RequestPayment is a mock implementation to satisfy the external.IPaymentService interface
func (m *MockPaymentService) RequestPayment(ctx context.Context, req external.PaymentRequest) (external.PaymentResponse, error) {
	return external.PaymentResponse{}, nil
}

func TestCreate(t *testing.T) {
	repo := &MockPaymentRepository{}
	service := &MockPaymentService{}
	gateway := NewPaymentGateway(service, repo)

	payment := &entity.Payment{
		ID:         "id1",
		OrderId:    "order1",
		Status:     entity.PaymentStatusPending,
		QrData:     "qr",
		ExternalId: "ext",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := gateway.Create(context.Background(), payment)
	assert.NoError(t, err)
	assert.True(t, repo.CreateCalled)
	assert.Equal(t, payment.ID, repo.CreateInput.ID)
}

func TestFindOne(t *testing.T) {
	repo := &MockPaymentRepository{
		FindOneOutput: &model.Payment{
			ID:         "id2",
			OrderId:    "order2",
			Status:     string(entity.PaymentStatusAuthorized),
			QrData:     "qr2",
			ExternalId: "ext2",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}
	service := &MockPaymentService{}
	gateway := NewPaymentGateway(service, repo)

	payment, err := gateway.FindOne(context.Background(), "id2")
	assert.NoError(t, err)
	assert.True(t, repo.FindOneCalled)
	assert.Equal(t, "id2", payment.ID)
	assert.Equal(t, entity.PaymentStatusAuthorized, payment.Status)
}

func TestUpdate(t *testing.T) {
	repo := &MockPaymentRepository{}
	service := &MockPaymentService{}
	gateway := NewPaymentGateway(service, repo)

	payment := &entity.Payment{
		ID:         "id3",
		OrderId:    "order3",
		Status:     entity.PaymentStatusNotAuthorized,
		QrData:     "qr3",
		ExternalId: "ext3",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := gateway.Update(context.Background(), payment)
	assert.NoError(t, err)
	assert.True(t, repo.UpdateCalled)
	assert.Equal(t, payment.ID, repo.UpdateInput.ID)
}

func TestFindByOrderId(t *testing.T) {
	repo := &MockPaymentRepository{
		FindByOrderIdOutput: &model.Payment{
			ID:         "id4",
			OrderId:    "order4",
			Status:     string(entity.PaymentStatusPending),
			QrData:     "qr4",
			ExternalId: "ext4",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}
	service := &MockPaymentService{}
	gateway := NewPaymentGateway(service, repo)

	payment, err := gateway.FindByOrderId(context.Background(), "order4")
	assert.NoError(t, err)
	assert.True(t, repo.FindByOrderIdCalled)
	assert.Equal(t, "order4", payment.OrderId)
	assert.Equal(t, entity.PaymentStatusPending, payment.Status)
}

func TestRequestPayment(t *testing.T) {
	repo := &MockPaymentRepository{}
	service := &MockPaymentService{}
	gateway := NewPaymentGateway(service, repo)

	payment := &entity.Payment{}
	metadata := dto.MetaData{}
	err := gateway.RequestPayment(context.Background(), payment, metadata)
	assert.NoError(t, err)
}
