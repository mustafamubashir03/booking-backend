'use strict';

import { DataTypes, QueryInterface } from "sequelize";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface: QueryInterface) {
    await queryInterface.addColumn('roomCategories', 'price', {
      type: DataTypes.DECIMAL,
      allowNull: false,
      defaultValue: 0
    });

    await queryInterface.removeColumn('rooms', 'price');
  },

  async down(queryInterface: QueryInterface) {
    await queryInterface.addColumn('rooms', 'price', {
      type: DataTypes.DECIMAL,
      allowNull: false,
      defaultValue: 0
    });

    await queryInterface.removeColumn('roomCategories', 'price');
  }
};