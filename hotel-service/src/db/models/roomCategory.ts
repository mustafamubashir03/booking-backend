import {
    CreationOptional,
    DataTypes,
    InferAttributes,
    InferCreationAttributes,
    Model,
} from 'sequelize';
import sequelize from './sequelize';


export enum RoomType {
    SINGLE = 'SINGLE',
    DOUBLE = 'DOUBLE',
    TRIPLE = 'TRIPLE',
    DELUXE = 'DELUXE',
    SUITE = 'SUITE',
    EXECUTIVE = 'EXECUTIVE',
    PENTHOUSE = 'PENTHOUSE',

}

class RoomCategory extends Model<InferAttributes<RoomCategory>, InferCreationAttributes<RoomCategory>> {
    declare id: CreationOptional<number>;
    declare hotelId: number;
    declare roomType: RoomType;
    declare roomCount: number;
    declare price: number;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deletedAt: CreationOptional<Date | null>;
}

RoomCategory.init(
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
        price: {
            type: DataTypes.NUMBER,
            allowNull: false
        },
        roomType: {
            type: DataTypes.ENUM(...Object.values(RoomType)),
            allowNull: false,
        },
        roomCount: {
            type: DataTypes.INTEGER,
            allowNull: false,
        },
        createdAt: DataTypes.DATE,
        updatedAt: DataTypes.DATE,
        deletedAt: {
            type: DataTypes.DATE,
            allowNull: true,
            defaultValue: null,
        },
    },
    {
        sequelize: sequelize,
        tableName: 'roomCategories',
        underscored: false,
        timestamps: true,
    },
);

export default RoomCategory;
