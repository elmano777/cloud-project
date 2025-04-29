import { relations } from 'drizzle-orm';
import { pgTable, serial, varchar } from 'drizzle-orm/pg-core';
import { cursosTable } from './cursos';

export const profesoresTable = pgTable('profesores', {
  id: serial('id').primaryKey(),
  nombre: varchar('nombre', { length: 255 }).notNull(),
  asignatura: varchar('asignatura', { length: 255 }).notNull(),
  email: varchar('email', { length: 255 }).notNull().unique(),
});

export const profesoresRelations = relations(profesoresTable, ({ many }) => ({
  cursos: many(cursosTable),
}));

export type Profesor = typeof profesoresTable.$inferSelect;
export type NewProfesor = typeof profesoresTable.$inferInsert;
