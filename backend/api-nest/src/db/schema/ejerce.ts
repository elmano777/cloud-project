import { pgTable, serial, integer, timestamp } from 'drizzle-orm/pg-core';
import { relations } from 'drizzle-orm';
import { usuariosTable } from './usuario';

export const ejerceTable = pgTable('ejerce', {
  id: serial('id').primaryKey(),
  profesor_id: integer('profesor_id')
    .notNull()
    .references(() => usuariosTable.id),
  curso_id: integer('curso_id').notNull(), // ID del curso del microservicio externo
  fecha_asignacion: timestamp('fecha_asignacion').defaultNow(),
});

export const ejerceRelations = relations(ejerceTable, ({ one }) => ({
  profesor: one(usuariosTable, {
    fields: [ejerceTable.profesor_id],
    references: [usuariosTable.id],
  }),
}));
