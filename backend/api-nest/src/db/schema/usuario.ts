import {
  pgTable,
  serial,
  varchar,
  boolean,
  timestamp,
} from 'drizzle-orm/pg-core';
import { relations } from 'drizzle-orm';
import { ejerceTable } from './ejerce';

export const usuariosTable = pgTable('usuarios', {
  id: serial('id').primaryKey(),
  nombre: varchar('nombre', { length: 100 }).notNull(),
  apellido: varchar('apellido', { length: 100 }).notNull(),
  email: varchar('email', { length: 100 }).notNull().unique(),
  password: varchar('password', { length: 255 }).notNull(),
  rol: varchar('rol', { length: 20 }).notNull().default('estudiante'), // Puede ser 'profesor' o 'estudiante'
  activo: boolean('activo').default(true),
  fechaCreacion: timestamp('fecha_creacion').defaultNow(),
  fechaActualizacion: timestamp('fecha_actualizacion').defaultNow(),
});

export const usuariosRelations = relations(usuariosTable, ({ many }) => ({
  ejerce: many(ejerceTable),
}));
