import { Injectable, NotFoundException } from '@nestjs/common';
import { db } from '../db/index';
import { ejerceTable } from '../db/schema/ejerce';
import { eq } from 'drizzle-orm';
import { CreateEjerceDto } from './dto/create-ejerce.dto';
import { UpdateEjerceDto } from './dto/update-ejerce.dto';

@Injectable()
export class EjerceService {
  async create(createEjerceDto: CreateEjerceDto) {
    const [nuevo] = await db
      .insert(ejerceTable)
      .values({
        ...createEjerceDto,
        fecha_asignacion: createEjerceDto.fecha_asignacion
          ? new Date(createEjerceDto.fecha_asignacion)
          : new Date(),
      })
      .returning();
    return nuevo;
  }

  async findAll() {
    return await db.select().from(ejerceTable);
  }

  async findOne(id: number) {
    const ejerce = await db
      .select()
      .from(ejerceTable)
      .where(eq(ejerceTable.id, id));
    if (ejerce.length === 0) {
      throw new NotFoundException(`Registro ejerce con ID ${id} no encontrado`);
    }
    return ejerce[0];
  }

  async update(id: number, updateEjerceDto: UpdateEjerceDto) {
    // Verificar si existe
    await this.findOne(id);

    const actualizado = await db
      .update(ejerceTable)
      .set({
        ...updateEjerceDto,
        fecha_asignacion: updateEjerceDto.fecha_asignacion
          ? new Date(updateEjerceDto.fecha_asignacion)
          : undefined,
      })
      .where(eq(ejerceTable.id, id))
      .returning();

    if (actualizado.length === 0) {
      throw new NotFoundException(`Registro ejerce con ID ${id} no encontrado`);
    }
    return actualizado[0];
  }

  async remove(id: number) {
    // Verificar si existe
    await this.findOne(id);

    const eliminado = await db
      .delete(ejerceTable)
      .where(eq(ejerceTable.id, id))
      .returning();

    return {
      message: 'Registro ejerce eliminado correctamente',
      ejerceEliminado: eliminado[0].id,
    };
  }
}
