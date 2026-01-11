from fastapi import FastAPI

app = FastAPI(title="ML API")

@app.get("/api/health")
async def root():
    return {"message":"Servir is running"}