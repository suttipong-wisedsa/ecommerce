"use client";

import { Card, Button } from "antd";

const { Meta } = Card;

export default function ProductCard({ product, onAdd }: any) {
  console.log(product.photo)
  return (
    <Card
      hoverable
      cover={<img src="product.photo" />}
    >
      <Meta title={product?.name} description={`${product?.price} บาท`} />

      <Button
        type="primary"
        style={{ marginTop: 10 }}
        block
        onClick={() => onAdd(product)}
      >
        Add to Cart
      </Button>
    </Card>
  );
}