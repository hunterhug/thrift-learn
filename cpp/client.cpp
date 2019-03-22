#include <iostream>

// lib
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/transport/TSocket.h>
#include <thrift/transport/TTransportUtils.h>
#include <boost/shared_ptr.hpp>
using namespace std;
using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;
using boost::shared_ptr;

// project
#include "gen-cpp/timeServe.h"


int main() {
  std::shared_ptr<TTransport> socket(new TSocket("localhost", 9090));
  std::shared_ptr<TTransport> transport(new TBufferedTransport(socket));
  std::shared_ptr<TProtocol> protocol(new TBinaryProtocol(transport));

  timeServeClient client(protocol);

    // open connect
  transport->open();
//  client.ping();
  //std::cout << "ping()" << std::endl;
  auto timeNow = client.getCurrtentTime();
  std::cout << timeNow << std::endl;
  transport->close();
  return 0;
}
