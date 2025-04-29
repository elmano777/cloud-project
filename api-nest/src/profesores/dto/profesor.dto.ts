import { IsString, IsEmail, IsOptional } from '@nestjs/class-validator';

export class CreateProfesorDto {
  @IsString()
  nombre: string;

  @IsString()
  asignatura: string;

  @IsEmail({}, { message: 'Debe proporcionar un email válido' })
  email: string;
}

export class UpdateProfesorDto {
  @IsString()
  @IsOptional()
  nombre?: string;

  @IsString()
  @IsOptional()
  asignatura?: string;

  @IsEmail({}, { message: 'Debe proporcionar un email válido' })
  @IsOptional()
  email?: string;
}
