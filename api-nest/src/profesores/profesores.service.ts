import { Injectable } from '@nestjs/common';
import { drizzle } from 'drizzle-orm/node-postgres';
import { eq, like } from 'drizzle-orm';
import { profesoresTable } from '../db/schema/profesores';
import { cursosTable } from '../db/schema/cursos';
import { Pool } from 'pg';
import { CreateProfesorDto, UpdateProfesorDto } from './dto/profesor.dto';
import { Profesor } from '../db/schema/profesores';

@Injectable()
export class ProfesoresService {
  private db: ReturnType<typeof drizzle>;

  constructor() {
    // Aquí podrías inyectar la configuración desde el módulo
    const pool = new Pool({
      connectionString: process.env.DATABASE_URL,
    });
    this.db = drizzle(pool);
  }

  // CREATE - Crear un nuevo profesor
  async create(createProfesorDto: CreateProfesorDto): Promise<Profesor> {
    try {
      const result = await this.db
        .insert(profesoresTable)
        .values(createProfesorDto)
        .returning();
      return result[0];
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Unknown error';
      throw new Error(`Error al crear profesor: ${message}`);
    }
  }

  // READ - Obtener todos los profesores
  async findAll(): Promise<Profesor[]> {
    try {
      return await this.db.select().from(profesoresTable);
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Unknown error';
      throw new Error(`Error al obtener profesores: ${message}`);
    }
  }

  // READ - Obtener un profesor por ID
  async findOne(id: number): Promise<Profesor> {
    try {
      const result = await this.db
        .select()
        .from(profesoresTable)
        .where(eq(profesoresTable.id, id));

      if (result.length === 0) {
        throw new Error(`Profesor con ID ${id} no encontrado`);
      }

      return result[0];
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Unknown error';
      throw new Error(`Error al buscar profesor: ${message}`);
    }
  }

  // READ - Buscar profesores por nombre
  async findByName(nombre: string): Promise<Profesor[]> {
    try {
      return await this.db
        .select()
        .from(profesoresTable)
        .where(like(profesoresTable.nombre, `%${nombre}%`));
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Unknown error';
      throw new Error(`Error al buscar profesores por nombre: ${message}`);
    }
  }

  // READ - Buscar profesores por asignatura
  async findByAsignatura(asignatura: string): Promise<Profesor[]> {
    try {
      return await this.db
        .select()
        .from(profesoresTable)
        .where(like(profesoresTable.asignatura, `%${asignatura}%`));
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Unknown error';
      throw new Error(`Error al buscar profesores por asignatura: ${message}`);
    }
  }

  // READ - Obtener profesor con sus cursos
  async findWithCursos(
    id: number,
  ): Promise<{ profesor: Profesor; cursos: any[] }> {
    try {
      const profesor = await this.findOne(id);

      const cursos = await this.db
        .select()
        .from(cursosTable)
        .where(eq(cursosTable.id_profesor, id));

      return { profesor, cursos };
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Unknown error';
      throw new Error(`Error al obtener profesor con curso: ${message}`);
    }
  }

  // UPDATE - Actualizar un profesor
  async update(
    id: number,
    updateProfesorDto: UpdateProfesorDto,
  ): Promise<Profesor> {
    try {
      const result = await this.db
        .update(profesoresTable)
        .set(updateProfesorDto)
        .where(eq(profesoresTable.id, id))
        .returning();

      if (result.length === 0) {
        throw new Error(`Profesor con ID ${id} no encontrado`);
      }

      return result[0];
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Unknown error';
      throw new Error(`Error al actualizar profesor: ${message}`);
    }
  }

  // DELETE - Eliminar un profesor
  async remove(id: number): Promise<Profesor> {
    try {
      const result = await this.db
        .delete(profesoresTable)
        .where(eq(profesoresTable.id, id))
        .returning();

      if (result.length === 0) {
        throw new Error(`Profesor con ID ${id} no encontrado`);
      }

      return result[0];
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Unknown error';
      throw new Error(`Error al eliminar profesor ${message}`);
    }
  }
}
