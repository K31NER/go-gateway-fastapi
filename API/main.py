from fastapi import FastAPI
from server import lifespan
from api.routes import predict

app = FastAPI(title="ML API", lifespan=lifespan)

@app.get("/api/health")
async def root():
    return {"message":"API ML is running"}

app.include_router(predict.router)