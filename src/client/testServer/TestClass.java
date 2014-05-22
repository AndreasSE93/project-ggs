package testServer;

import java.io.IOException;

import org.json.JSONException;

import clientNetworking.Connection;
import clientNetworking.NetManager;

import packageManaging.CreateGameMessage;
import packageManaging.CreateGameMessageEncoder;
import packageManaging.TiarStartedMessage;
import packageManaging.TiarStartedMessageEncoder;

public class TestClass implements Runnable{
	

	/**
	 * @param args
	 * 
	 * 
	 */
	
		TestClass(){
			
			
		}
		
	public void init(){
		Connection conn = new Connection("130.243.137.96", 8080);
		NetManager network = new NetManager(conn);
		CreateGameMessageEncoder gme = new CreateGameMessageEncoder();
		TiarStartedMessageEncoder tiStarter = new TiarStartedMessageEncoder();
		try {
			network.connectToServer();
		} catch (IOException e2) {
			// TODO Auto-generated catch block
			e2.printStackTrace();
		}
		
		String JSONtext = "";
		try {
			JSONtext = gme.encode(new CreateGameMessage("Achtung Die Kurve", 1, "Viktor"));
			sendMessage(JSONtext,network);
			JSONtext = gme.encode(new CreateGameMessage("Achtung Die Kurve", 1, "Viktor"));
			sendMessage(JSONtext,network);
		} catch (JSONException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		

		
		try {
			JSONtext = tiStarter.encode(new TiarStartedMessage());
			sendMessage(JSONtext,network);
			System.out.println("Sent message");
		} catch (JSONException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
	

		
		while (true) {
			
			try {
				String mess = receiveMessage(network);
				System.out.println(Thread.currentThread().getId() + " " + mess);
			} catch (IOException e1) {
				e1.printStackTrace();
			}
		}
		
		
	}
	
	public void init2(){
		Connection conn = new Connection("130.243.137.96", 8080);
		NetManager network = new NetManager(conn);
		CreateGameMessageEncoder gme = new CreateGameMessageEncoder();
		TiarStartedMessageEncoder tiStarter = new TiarStartedMessageEncoder();
		try {
			network.connectToServer();
		} catch (IOException e2) {
			// TODO Auto-generated catch block
			e2.printStackTrace();
		}
		
		String JSONtext = "";
		try {
			JSONtext = gme.encode(new CreateGameMessage("TicTacToe", 1, "Viktor"));
			sendMessage(JSONtext,network);
			JSONtext = gme.encode(new CreateGameMessage("TicTacToe", 1, "Viktor"));
			sendMessage(JSONtext,network);
		} catch (JSONException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		

		
		try {
			JSONtext = tiStarter.encode(new TiarStartedMessage());
			sendMessage(JSONtext,network);
		
		} catch (JSONException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
	

		
		while (true) {
			
			try {
				String mess = receiveMessage(network);
				System.out.println(Thread.currentThread().getId() + " " + mess);
			} catch (IOException e1) {
				e1.printStackTrace();
			}
		}
		
	}
	
	public void sendMessage(String message, NetManager network) {
		network.send(message);
	}

	public String receiveMessage(NetManager network) throws IOException {
		return network.receiveMessage();
		

	}

	public static void main(String[] args) {
		
		//Init connection count
		for(int i =0; i< 10000; i++){
		(new Thread(new TestClass())).start();
		}

		
	}

	@Override
	public void run() {

		if(Thread.currentThread().getId() % 2 == 0)
		init();
		else
		init2();
		
	}

}
