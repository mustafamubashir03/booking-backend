'use strict';

import { QueryInterface } from "sequelize";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface: QueryInterface) {
    await queryInterface.addConstraint('roomCategories', {
      fields: ['hotelId'],
      type: 'foreign key',
      name: 'fk_roomCategories_hotelId',
      references: {
        table: 'hotels',
        field: 'id',
      },
      onUpdate: 'CASCADE',
      onDelete: 'SET NULL',
    });
  },

  async down(queryInterface: QueryInterface) {
    await queryInterface.removeConstraint('roomCategories', 'fk_roomCategories_hotelId');
  },
};
