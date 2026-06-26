import {
    CreationOptional,
    DataTypes,
    InferAttributes,
    InferCreationAttributes,
    Model,
} from 'sequelize';
import sequelize from './sequelize';



class Room extends Model<InferAttributes<Room>, InferCreationAttributes<Room>> {
    declare id: CreationOptional<number>;
    declare hotelId: number;
    declare roomCategoryId: number;
    declare dateOfAvailability: Date;
    declare price: number;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deletedAt: CreationOptional<Date | null>;
    declare bookingId?: number | null
}

Room.init(
    {
        id: {
            type: DataTypes.INTEGER,
            autoIncrement: true,
            primaryKey: true,
        },
        hotelId: {
            type: DataTypes.INTEGER,
            allowNull: false,
        },
        roomCategoryId: {
            type: DataTypes.INTEGER,
            allowNull: false,
        },
        dateOfAvailability: {
            type: DataTypes.DATE,
            allowNull: false,
        },
        price: {
            type: DataTypes.NUMBER,
            allowNull: false
        },
        createdAt: DataTypes.DATE,
        updatedAt: DataTypes.DATE,
        deletedAt: {
            type: DataTypes.DATE,
            allowNull: true,
            defaultValue: null,
        },
        bookingId: {
            type: DataTypes.INTEGER,
            allowNull: true,
            defaultValue: null,
        }
    },
    {
        sequelize: sequelize,
        tableName: 'room',
        underscored: true,
        timestamps: true,
    },
);

export default Room;
