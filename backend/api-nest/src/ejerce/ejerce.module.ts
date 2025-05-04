import { Module } from '@nestjs/common';
import { EjerceController } from './ejerce.controller';
import { EjerceService } from './ejerce.service';

@Module({
  controllers: [EjerceController],
  providers: [EjerceService],
})
export class EjerceModule {}
