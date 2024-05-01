CREATE TABLE fundsusd (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    ticker text NOT NULL,
    amount decimal NOT NULL,
    pricePerItem decimal NOT NULL,
    purchasePrice decimal NOT NULL,
    priceCurrent decimal NOT NULL,
    percentChanges decimal NOT NULL,
    yearlyInvestment decimal NOT NULL,
    clearMoney decimal NOT NULL,
    datePurchase time NOT NULL,
    dateLastUpdate time NOT NULL,
    type text NOT NULL
);