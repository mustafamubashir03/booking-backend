'use strict';
import { QueryInterface } from "sequelize";
/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface: QueryInterface, Sequelize: any) {
    await queryInterface.sequelize.query(`
      CREATE TABLE IF NOT EXISTS roomCategories(
        id INT AUTO_INCREMENT PRIMARY KEY,
        hotelId INT NOT NULL,
        roomType ENUM('SINGLE', 'DOUBLE', 'TRIPLE', 'DELUXE', 'SUITE', 'EXECUTIVE', 'PENTHOUSE') NOT NULL,
        roomCount INT NOT NULL,
        createdAt DATE,
        updatedAt DATE,
        deletedAt DATE
      )
      `)
  },

  async down(queryInterface: QueryInterface, Sequelize: any) {
    await queryInterface.sequelize.query(`
      DROP TABLE IF EXISTS roomCategories;
      `)
  }
};
