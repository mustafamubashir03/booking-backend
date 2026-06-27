import { CreationAttributes, Model, ModelStatic, WhereOptions } from "sequelize";

export class BaseRepository<T extends Model> {
    protected model: ModelStatic<T>

    constructor(model: ModelStatic<T>) {
        this.model = model
    }
    async findById(id: number): Promise<T | null> {
        const record = await this.model.findByPk(id)
        if (!record) {
            return null
        }
        return record

    }
    async findAll(): Promise<T[]> {
        const records = await this.model.findAll()
        if (!records) {
            return []
        }
        return records

    }
    async deleteById(whereOptions: WhereOptions<T>): Promise<void> {
        const record = await this.model.destroy({
            where: {
                ...whereOptions
            }
        })
        if (!record) {
            throw new Error("Record not found")
        }

        return;
    }
    async create(data: CreationAttributes<T>): Promise<T> {
        const record = await this.model.create(data)
        return record
    }
    async update(id: number, data: Partial<T>): Promise<T | null> {
        const [affectedCount, affectedRows] = await this.model.update(data, {
            where: { id },
            returning: true,
        } as any);
        if (affectedCount === 0 || !affectedRows || affectedRows.length === 0) {
            return null;
        }
        return affectedRows[0] as T;
    }

}