import { pgTable, serial, varchar, text, integer } from 'drizzle-orm/pg-core';
import { profesoresTable } from './profesores';
import { relations } from 'drizzle-orm';

export const cursosTable = pgTable('cursos', {
  id: serial('id').primaryKey(),
  nombre_curso: varchar('nombre_curso', { length: 255 }).notNull(),
  descripcion: text('descripcion'),
  id_profesor: integer('id_profesor')
    .notNull()
    .references(() => profesoresTable.id),
});

export const cursosRelations = relations(cursosTable, ({ one }) => ({
  profesor: one(profesoresTable, {
    fields: [cursosTable.id_profesor],
    references: [profesoresTable.id],
  }),
}));
