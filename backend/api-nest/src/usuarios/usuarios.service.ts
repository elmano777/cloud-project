import { Injectable, NotFoundException } from '@nestjs/common';
import { CreateUsuarioDto } from './dto/create-usuario.dto';
import { UpdateUsuarioDto } from './dto/update-usuario.dto';
import { db } from '../db';
import { usuariosTable } from '../db/schema/usuario';
import { eq } from 'drizzle-orm';
import * as bcrypt from 'bcrypt';

@Injectable()
export class UsuariosService {
  async create(createUsuarioDto: CreateUsuarioDto) {
    // Encriptar la contraseña
    const hashedPassword = await bcrypt.hash(createUsuarioDto.password, 10);

    // Crear el usuario con la contraseña encriptada
    const nuevoUsuario = await db
      .insert(usuariosTable)
      .values({
        ...createUsuarioDto,
        password: hashedPassword,
      })
      .returning();

    return nuevoUsuario[0];
  }

  async findAll(limit = 10, page = 1) {
    const offset = (page - 1) * limit;
    return await db.select().from(usuariosTable).limit(limit).offset(offset);
  }

  async findOne(id: number) {
    const usuario = await db
      .select()
      .from(usuariosTable)
      .where(eq(usuariosTable.id, id));

    if (usuario.length === 0) {
      throw new NotFoundException(`Usuario con ID ${id} no encontrado`);
    }

    return usuario[0];
  }

  async update(id: number, updateUsuarioDto: UpdateUsuarioDto) {
    // Verificar si el usuario existe
    await this.findOne(id);

    // Si se actualiza la contraseña, encriptarla
    if (updateUsuarioDto.password) {
      updateUsuarioDto.password = await bcrypt.hash(
        updateUsuarioDto.password,
        10,
      );
    }

    // Actualizar el usuario
    const usuarioActualizado = await db
      .update(usuariosTable)
      .set({
        ...updateUsuarioDto,
        fechaActualizacion: new Date(),
      })
      .where(eq(usuariosTable.id, id))
      .returning();

    return usuarioActualizado[0];
  }

  async login(email: string, password: string) {
    console.log(`Intentando login con email: ${email}`);

    // Buscar usuario por email
    const usuarios = await db
      .select()
      .from(usuariosTable)
      .where(eq(usuariosTable.email, email));

    console.log('Usuarios encontrados:', usuarios.length);

    const usuario = usuarios[0];
    if (!usuario) {
      console.log('No se encontró el usuario');
      throw new NotFoundException('Credenciales incorrectas');
    }

    console.log('Usuario encontrado:', usuario.email);
    console.log('Contraseña proporcionada:', password);
    console.log('Contraseña almacenada (hash):', usuario.password);

    // Comparar contraseñas
    const passwordMatch = await bcrypt.compare(password, usuario.password);
    console.log('¿Contraseña coincide?:', passwordMatch);

    if (!passwordMatch) {
      console.log('La contraseña no coincide');
      throw new NotFoundException('Credenciales incorrectas');
    }

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const { password: _, ...usuarioSinPassword } = usuario;
    return usuarioSinPassword;
  }
  async remove(id: number) {
    // Verificar si el usuario existe
    await this.findOne(id);

    // Eliminar el usuario
    const usuarioEliminado = await db
      .delete(usuariosTable)
      .where(eq(usuariosTable.id, id))
      .returning();

    return {
      message: 'Usuario eliminado correctamente',
      usuarioEliminado: usuarioEliminado[0].email,
    };
  }
}
