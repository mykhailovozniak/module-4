package module_4

import "fmt"

type Builder interface {
	AddWebsite(w string)
	AddAndroidApp(a string)
	AddiOSApp(i string)
}

type Developer struct {
	builder Builder
}

type Product struct {
	features string
}

type ReleaseBuilder struct {
	product *Product
}

func (b *ReleaseBuilder) AddWebsite(w string)  {
	b.product.features += "We just implemented a new website " + w + " name \n"
}

func (b *ReleaseBuilder) AddAndroidApp(a string)  {
	b.product.features += "We just implemented a new Android App " + a + " name \n"
}

func (b *ReleaseBuilder) AddiOSApp(i string)  {
	b.product.features += "We just implemented a new iOS App " + i + " name \n"
}

func (p *Product) Show() string {
	return p.features
}

func (p *Product) Reset() {
	p.features = ""
}

func (d *Developer) buildMinimalViableProduct()  {
	d.builder.AddWebsite("test.com")
}

func (d *Developer) buildFullFeaturedProduct() {
	d.builder.AddWebsite("test.com")
	d.builder.AddAndroidApp("android app")
	d.builder.AddiOSApp(" iOS app")
}

func builderClientCode() {
	product := new(Product)
	developer := Developer{&ReleaseBuilder{product}}

	developer.buildMinimalViableProduct()
	fmt.Println(product.Show())

	product.Reset()

	developer.buildFullFeaturedProduct()
	fmt.Println(product.Show())
}
