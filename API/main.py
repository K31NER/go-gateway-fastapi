from routes import predict
from fastapi import FastAPI
from server import lifespan

app = FastAPI(title="ML API", lifespan=lifespan)

@app.get("/api/health")
async def root():
    return {"message":"Servir is running"}

app.include_router(predict.router)