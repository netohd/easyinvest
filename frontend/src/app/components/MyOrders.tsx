import { Order } from "../models"

// Server Components (antes a página inteira era renderizada, agora em nível de comp.)
async function getOrder(wallet_id: string): Promise<Order[]> {
  const response = await fetch(`http://localhost:8000/wallets/${wallet_id}/orders`);
  return response.json();
}

export default async function MyOrders(props: { wallet_id: string }) {
  const orders = await getOrder(props.wallet_id)

  return (
      <ul>
        {orders.map((order) => (
          <li key={order.id}>
            {order.Asset.id} - {order.shares} - R$ {order.price} - {order.status}
          </li>
        ))}
      </ul>
  )
}