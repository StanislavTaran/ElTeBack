CREATE TABLE `bank` (
                         `id` int AUTO_INCREMENT PRIMARY KEY,
                         `name` varchar(255),
                         `interestRate` INT,
                         `maxLoan` INT,
                         `minDownPayment` INT,
                         `loanTerm` INT
);

ALTER TABLE bank AUTO_INCREMENT=1;
