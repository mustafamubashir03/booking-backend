import {
    CreationOptional,
    DataTypes,
    InferAttributes,
    InferCreationAttributes,
    Model,
} from 'sequelize';
import sequelize from './sequelize';


enum RoomType {
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
        roomType: {
            type: DataTypes.ENUM,
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
        tableName: 'roomCategory',
        underscored: true,
        timestamps: true,
    },
);

export default RoomCategory;
