export type Item = {
  shortDescription: string
  price: string // $0.00
}

export type Receipt = {
  retailer: string
  purchaseDate: string // YYYY-MM-DD
  purchaseTime: string // HH:MM (24-hour)
  total: string // $0.00
  items: Item[]
  points?: number
  id?: string
}
