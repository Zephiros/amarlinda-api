# amarlinda-back

Entidades

*User*
- Id
- Name
- Email
- Password
- CreatedAt
- UpdatedAt

*Product*
- Id
- Code
- Name
- Description
- AmountAvailable
- Price
- PricePromotion
- Active
- Promotion
- CreatedAt
- UpdatedAt

*Client*
- Id
- Name
- Phone
- CreatedAt
- UpdatedAt

*Payment*
- Id
- Name
- CreatedAt
- UpdatedAt

**Order**
- Id
- PaymentId
- SubTotalValue
- TotalValue
- Observation
- CreatedAt
- UpdatedAt

*OrderItem*
- Id
- OrderId
- ProductId
- Promotion
- Quantity
- Price
- CreatedAt
- UpdatedAt

*OrderIn : Order*
- Supplier

*OrderOut : Order*
- ClientId
