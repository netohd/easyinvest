import MyOrders from "@/app/components/MyOrders"

export default async function HomeBrokerPage({
  params
}: {
  params: { wallet_id: string, asset_id: string }
}) {

  return (
    <div>
      <h1>Home Broker</h1>
      {/* Tailwind usa classes utilitárias, utility first */}
      <div className="flex flex-row">
        <div className="flex flex-col">
          <div>formulario</div>
          <div>
            <MyOrders wallet_id={params.wallet_id} />
          </div>
        </div>

        <div>
          gráfico
        </div>
      </div>
    </div>
  )
}