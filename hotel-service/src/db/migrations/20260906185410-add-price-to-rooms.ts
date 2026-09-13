'use strict';

import { QueryInterface, DataTypes } from "sequelize";

module.exports = {
  async up(queryInterface: QueryInterface) {
    await queryInterface.addColumn('rooms', 'price', {
      type: DataTypes.INTEGER,
      allowNull: false,
      defaultValue: 0,
    });
  },

  async down(queryInterface: QueryInterface) {
    await queryInterface.removeColumn('rooms', 'price');
  }
};