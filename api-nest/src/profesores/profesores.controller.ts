import {
  Body,
  Controller,
  Get,
  HttpException,
  HttpStatus,
  Post,
  Query,
} from '@nestjs/common';
import { ProfesoresService } from './profesores.service';
import { CreateProfesorDto } from './dto/profesor.dto';

@Controller('profesores')
export class ProfesoresController {
  constructor(private readonly profesoresService: ProfesoresService) {}

  @Post()
  async create(@Body() createProfesorDto: CreateProfesorDto) {
    try {
      return await this.profesoresService.create(createProfesorDto);
    } catch (error) {
      const errorMessage =
        error instanceof Error ? error.message : 'An unexpected error occurred';
      throw new HttpException(errorMessage, HttpStatus.BAD_REQUEST);
    }
  }

  @Get()
  async findAll(
    @Query('nombre') nombre?: string,
    @Query('asignatura') asignatura?: string,
  ) {
    try {
      if (nombre) {
        return await this.profesoresService.findByName(nombre);
      }

      if (asignatura) {
        return await this.profesoresService.findByAsignatura(asignatura);
      }

      return await this.profesoresService.findAll();
    } catch (error) {
      const errorMessage =
        error instanceof Error ? error.message : 'An unexpected error occurred';
      throw new HttpException(errorMessage, HttpStatus.INTERNAL_SERVER_ERROR);
    }
  }
}
