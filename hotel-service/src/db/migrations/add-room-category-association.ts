'use strict';

import { QueryInterface } from "sequelize";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface: QueryInterface) {
    await queryInterface.addConstraint('rooms', {
      fields: ['roomCategoryId'],
      type: 'foreign key',
      name: 'fk_rooms_roomCategoryId',
      references: {
        table: 'roomCategories',
        field: 'id',
      },
      onUpdate: 'CASCADE',
      onDelete: 'SET NULL',
    });
  },

  async down(queryInterface: QueryInterface) {
    await queryInterface.removeConstraint('rooms', 'fk_rooms_roomCategoryId');
  },
};
