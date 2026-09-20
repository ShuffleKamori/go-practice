from fastapi import FastAPI, HTTPException
from sqlmodel import Field, Session, SQLModel, create_engine, select

app = FastAPI(title="Bicycle Store API")

engine = create_engine("sqlite:///store.db", connect_args={"check_same_thread": False})


class Product(SQLModel, table=True):
    id: int | None = Field(default=None, primary_key=True)
    name: str
    price: int
    image: str | None = None
    description: str | None = None


class ProductCreate(SQLModel):
    name: str
    price: int
    image: str | None = None
    description: str | None = None


SEED = [
    {"name": "Stels Navigator 21", "price": 24990, "image": "BuyCard/Recycle.jpg", "description": "Горный велосипед"},
    {"name": "Aist On-Trail", "price": 19990, "image": "BuyCard/Recycle.jpg", "description": "Городской хардтейл"},
    {"name": "Forward Apache", "price": 15990, "image": "BuyCard/Recycle.jpg", "description": "Легкий стартовый байк"},
    {"name": "Author Aventon", "price": 27990, "image": "BuyCard/Recycle.jpg", "description": "Комфортный универсал"},
    {"name": "Merida Matts", "price": 38990, "image": "BuyCard/Recycle.jpg", "description": "Полупрофессиональный CT"},
    {"name": "Titan Skipper", "price": 9990, "image": "BuyCard/Recycle.jpg", "description": "Детский велосипед"},
    {"name": "Cube Aim EX", "price": 45990, "image": "BuyCard/Recycle.jpg", "description": "Хардтейл на 29 колесах"},
    {"name": "Nordway Crossline", "price": 18990, "image": "BuyCard/Recycle.jpg", "description": "Городской с багажником"},
    {"name": "GT Aggressor Pro", "price": 29990, "image": "BuyCard/Recycle.jpg", "description": "Спортивный стандарт"},
    {"name": "Jamis Trail X", "price": 21990, "image": "BuyCard/Recycle.jpg", "description": "Маневренный трейловый"},
    {"name": "Scout Enduro", "price": 32990, "image": "BuyCard/Recycle.jpg", "description": "Эндуро для гор"},
    {"name": "Giant Talon 3", "price": 37990, "image": "BuyCard/Recycle.jpg", "description": "Надежный хардтейл"},
    {"name": "Trek Marlin 5", "price": 42990, "image": "BuyCard/Recycle.jpg", "description": "Универсальный кроссовый"},
    {"name": "Specialized Rockhopper", "price": 47990, "image": "BuyCard/Recycle.jpg", "description": "Премиальный хардтейл"},
    {"name": "Cannondale Trail 7", "price": 39990, "image": "BuyCard/Recycle.jpg", "description": "Сбалансированный трейл"},
    {"name": "Ghost Kato FS", "price": 56990, "image": "BuyCard/Recycle.jpg", "description": "Двухподвес для профи"},
]


def create_db_and_seed() -> None:
    SQLModel.metadata.create_all(engine)
    with Session(engine) as s:
        if s.exec(select(Product)).first() is None:
            s.add_all(Product(**p) for p in SEED)
            s.commit()


create_db_and_seed()


@app.get("/products", response_model=list[Product])
def get_products():
    with Session(engine) as s:
        return list(s.exec(select(Product)))


@app.get("/products/{product_id}", response_model=Product)
def get_product(product_id: int):
    with Session(engine) as s:
        p = s.get(Product, product_id)
        if p is None:
            raise HTTPException(404, "product not found")
        return p


@app.post("/products", response_model=Product, status_code=201)
def create_product(body: ProductCreate):
    with Session(engine) as s:
        p = Product(**body.model_dump())
        s.add(p)
        s.commit()
        s.refresh(p)
        return p


@app.put("/products/{product_id}", response_model=Product)
def update_product(product_id: int, body: ProductCreate):
    with Session(engine) as s:
        p = s.get(Product, product_id)
        if p is None:
            raise HTTPException(404, "product not found")
        for key, value in body.model_dump().items():
            setattr(p, key, value)
        s.add(p)
        s.commit()
        s.refresh(p)
        return p


@app.delete("/products/{product_id}", status_code=204)
def delete_product(product_id: int):
    with Session(engine) as s:
        p = s.get(Product, product_id)
        if p is None:
            raise HTTPException(404, "product not found")
        s.delete(p)
        s.commit()