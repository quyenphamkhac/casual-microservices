import { Controller, Get } from '@nestjs/common';
import { EventPattern } from '@nestjs/microservices';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  async fakeCreateProduct(): Promise<any> {
    const product = {
      id: 1,
      name: 'Play Station 5',
    };
    await this.appService.publishEvent('product_created', product);
    return product;
  }

  @EventPattern('product_created')
  async handleProductCreatedEvent(data: Record<string, unknown>) {
    console.log('hello data', data);
  }
}
