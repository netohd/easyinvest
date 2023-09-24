import MyWallet from "../components/MyWallet";
import { WalletAsset } from "../models"

async function getWalletAssets(wallet_id: string): Promise<WalletAsset[]> {
  const response = await fetch(`http://localhost:8000/wallets/${wallet_id}/assets`);
  return response.json();
}

export default async function HomePage({
  params
}: {
  params: { wallet_id: string }
}) {
  const walletAssets = await getWalletAssets(params.wallet_id)

  return (
    <div>
      <h1>Meus Investimentos</h1>
      <MyWallet wallet_id={params.wallet_id} />
    </div>
  )
}