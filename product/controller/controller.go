package controller

import (
	"context"
	"errors"
	"sync"

	"github.com/larien/product/product/drivers/log"

	"github.com/larien/product/product/entity"
)

// Controller represents the methods implemented by this Controller layer
type Controller interface {
	List(ctx context.Context, userID string) ([]*entity.Product, error)
}

var (
	errListProducts = errors.New("an error occurred when listing the products")
)

type (
	// Product represents the product repository methods used by this controller that
	// the injected dependency must implement
	Product interface {
		// List lists the available products.
		List() (products []entity.Product, err error)
	}

	// Discount represents the discount repository methods used by this controller that
	// the injected dependency must implement
	Discount interface {
		// Get sends the product ID and user ID and obtains a discount percentage
		Get(ctx context.Context, productID, userID string) (int64, error)
	}
)

type controller struct {
	Product  Product
	Discount Discount
}

// New creates a new instance of Product controller to make business logic decisions
func New(product Product, discount Discount) Controller {
	return &controller{
		Product:  product,
		Discount: discount,
	}
}

// TODO - if userID is not sent, should we even try to obtain the discount?
func (c *controller) List(ctx context.Context, userID string) ([]*entity.Product, error) {
	log.Debug(ctx, "listing products")
	products, err := c.Product.List()
	if err != nil {
		log.Error(ctx, err, errListProducts.Error())
		return nil, errListProducts
	}
	if products == nil {
		log.Info(ctx, "no product was found")
		return nil, nil
	}

	var p []*entity.Product
	for _, product := range products {
		c.applyDiscount(ctx, &product, userID)
		p = append(p, &product)
	}


	log.Debug(ctx, "products were obtained with success")
	return p, nil
}

func (c *controller) applyDiscount(ctx context.Context, product *entity.Product, userID string) {
	log.Debug(ctx, "obtaining discount percentage")
	percentage, err := c.Discount.Get(ctx, product.ID, userID)
	if err != nil {
		log.Error(ctx, err, "failed to apply discount")
		return
	}
	log.Debugf(ctx, "discount percentage to be applied: %d", percentage)
	product.Discount.Percentage = percentage
	product.Discount.ValueInCents = (product.PriceInCents * percentage) / 100
}
