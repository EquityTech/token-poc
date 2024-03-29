pragma solidity ^0.5.0;

import "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";

contract TestToken is ERC20 {
  string public name = "TestToken";
  string public symbol = "TES";
  uint8 public decimals = 2;
  uint public INITIAL_SUPPLY = 10000;

  constructor() public {
    _mint(msg.sender, INITIAL_SUPPLY);
    }
}
