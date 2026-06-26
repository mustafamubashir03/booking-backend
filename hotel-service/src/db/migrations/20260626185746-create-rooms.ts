'use strict';

import { QueryInterface } from "sequelize";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      CREATE TABLE IF NOT EXISTS rooms (
        id INT AUTO_INCREMENT PRIMARY KEY,
        hotelId INT NOT NULL,
        roomCategoryId INT NOT NULL,
        dateOfAvailability DATE NOT NULL,
        price NUMBER NOT NULL,
        createdAt DATE,
        updatedAt DATE,
        deletedAt DATE,
        bookingId INT
      )
      
      `)
  },

  async down(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      DROP TABLE IF EXISTS rooms;
      `)
  }
};
