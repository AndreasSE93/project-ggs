package packageManaging;

public class InitializeClientMessage {
		final int PacketID = 99;
		String userName;
		
		public InitializeClientMessage(String u){
			userName = u;
			
		}

}
