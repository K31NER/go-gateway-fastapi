from fastapi import FastAPI
from contextlib import asynccontextmanager

@asynccontextmanager
async def lifespan(app: FastAPI):
    """ Maneja el ciclo de vida del servidor """
    print("Iniciando servidor...")
    
    
    yield 
    
    print("Cerrando servidor....")