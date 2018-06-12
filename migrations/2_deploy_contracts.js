const Recover = artifacts.require("Recover");

module.exports = async (deployer) => {
  try {
    deployer.deploy(Recover)
      .then(() => Recover.deployed)
  } catch(err) {
    console.log('ERROR on deploy:',err);
  }

};
